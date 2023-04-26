import {
  CACHE_MANAGER,
  HttpException,
  HttpStatus,
  Inject,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { UsersService } from 'src/users/users.service';
import * as bcrypt from 'bcrypt';
import { CreateUserDto } from 'src/users/dto/create-user.dto';
import { Cache } from 'cache-manager';
import { jwtConstants } from './constants';

@Injectable()
export class AuthService {
  constructor(
    private usersService: UsersService,
    private jwtService: JwtService,
    @Inject(CACHE_MANAGER) private cacheManager: Cache,
  ) {}

  async loginUser(email: string, pass: string): Promise<any> {
    const user = await this.usersService.FindByEmail(email);

    const isMatchPass = await bcrypt.compare(pass, user.password);

    if (!isMatchPass) {
      throw new UnauthorizedException();
    }

    const payload = { id: user.id, email: user.email };

    const access_token = await this.jwtService.signAsync(payload);

    await this.cacheManager.set(access_token, access_token);

    return {
      id: user.id,
      access_token,
    };
  }

  async signUp(createUserDto: CreateUserDto): Promise<any> {
    const verifyUser = await this.usersService.FindByEmail(createUserDto.email);

    if (verifyUser) {
      throw new HttpException('user not available', HttpStatus.BAD_REQUEST);
    }

    const users = await this.usersService.add(createUserDto);

    const result = {
      id: users.id,
      first_name: users.first_name,
      last_name: users.last_name,
      email: users.email,
    };

    return result;
  }

  async refreshAuthentication(refreshToken: string): Promise<any> {
    await this.verifyRefreshToken(refreshToken);
    await this.checkAvailabilityToken(refreshToken);
    const decodeToken = this.jwtService.decode(refreshToken);

    const payload = {
      id: decodeToken['id'],
      email: decodeToken['email'],
    };

    return this.jwtService.signAsync(payload);
  }

  async verifyRefreshToken(refreshToken: string) {
    try {
      await this.jwtService.verifyAsync(refreshToken, {
        secret: jwtConstants.secret,
      });
    } catch (error) {
      console.log(error);
      throw new UnauthorizedException();
    }
  }

  async checkAvailabilityToken(token: string): Promise<any> {
    const result = await this.cacheManager.get(token);
    if (!result) {
      throw new HttpException(
        'refresh token not found',
        HttpStatus.BAD_REQUEST,
      );
    }
  }

  async logoutUser(refreshToken: string): Promise<any> {
    await this.checkAvailabilityToken(refreshToken);
    return this.cacheManager.del(refreshToken);
  }
}
