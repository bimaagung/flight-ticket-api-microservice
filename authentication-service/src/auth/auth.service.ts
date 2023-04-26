import {
  CACHE_MANAGER,
  HttpException,
  HttpStatus,
  Inject,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { UsersService } from 'src/users/users.service';
import * as bcrypt from 'bcrypt';
import { CreateUserDto } from 'src/users/dto/create-user.dto';
import { Cache } from 'cache-manager';
import { JwtTokenManagerService } from 'src/security/jwt-token-manager/jwt-token-manager.service';

@Injectable()
export class AuthService {
  constructor(
    private usersService: UsersService,
    @Inject(CACHE_MANAGER) private cacheManager: Cache,
    private jwtTokenManagerService: JwtTokenManagerService,
  ) {}

  async loginUser(email: string, pass: string): Promise<any> {
    const user = await this.usersService.FindByEmail(email);

    const isMatchPass = await bcrypt.compare(pass, user.password);

    if (!isMatchPass) {
      throw new UnauthorizedException();
    }

    const payload = { id: user.id, email: user.email };

    const accessToken = await this.jwtTokenManagerService.createAccessToken(
      payload,
    );
    const refreshToken = await this.jwtTokenManagerService.refreshAccessToken(
      payload,
    );

    await this.cacheManager.set(refreshToken, refreshToken);

    return {
      accessToken,
      refreshToken,
    };
  }

  async refreshAuthentication(refreshToken: string): Promise<any> {
    await this.jwtTokenManagerService.verifyRefreshToken(refreshToken);
    await this.checkAvailabilityToken(refreshToken);
    const decodeToken = this.jwtTokenManagerService.decodePayload(refreshToken);

    const payload = {
      id: decodeToken['id'],
      email: decodeToken['email'],
    };

    return this.jwtTokenManagerService.createAccessToken(payload);
  }

  async logoutUser(refreshToken: string): Promise<any> {
    await this.checkAvailabilityToken(refreshToken);
    return this.cacheManager.del(refreshToken);
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
}
