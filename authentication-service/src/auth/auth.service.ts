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
    await this.cacheManager.get(refreshToken);
    const { id, email } = this.decodePayload(refreshToken);
    return this.jwtService.signAsync({ id, email });
  }

  async verifyRefreshToken(refreshToken: string) {
    try {
      await this.jwtService.verifyAsync(refreshToken, {
        secret: jwtConstants.secret,
      });
    } catch (error) {
      throw new UnauthorizedException();
    }
  }

  decodePayload(refreshToken: string) {
    const artifacts = this.jwtService.decode(refreshToken);
    return artifacts['decoded'].payload;
  }
}
