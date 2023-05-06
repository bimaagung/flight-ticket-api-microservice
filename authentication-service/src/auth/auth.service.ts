import {
  CACHE_MANAGER,
  HttpException,
  HttpStatus,
  Inject,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { UsersService } from './../users/users.service';
import * as bcrypt from 'bcrypt';
import { UpdatePasswordDto } from 'src/auth/dto/update-password.dto';
import { Cache } from 'cache-manager';
import { JwtTokenManagerService } from './../security/jwt-token-manager/jwt-token-manager.service';
import { CreateUserDto } from './../users/dto/create-user.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { User } from './../users/entities/user.entity';
import { Repository } from 'typeorm';

@Injectable()
export class AuthService {
  constructor(
    @InjectRepository(User)
    private usersRepository: Repository<User>,
    private usersService: UsersService,
    private jwtTokenManagerService: JwtTokenManagerService,
    @Inject(CACHE_MANAGER)
    private cacheManager: Cache,
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
      throw new HttpException('refresh token not found', HttpStatus.NOT_FOUND);
    }
  }

  async signUp(createUserDto: CreateUserDto): Promise<any> {
    const verifyUser = await this.usersService.FindByEmail(createUserDto.email);

    if (verifyUser) {
      throw new HttpException('user not available', HttpStatus.NOT_FOUND);
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

  async updatePassword(
    idUser: string,
    updatePasswordDto: UpdatePasswordDto,
  ): Promise<any> {
    if (updatePasswordDto.new_password !== updatePasswordDto.retype_password) {
      throw new HttpException('passwords do not match', HttpStatus.BAD_REQUEST);
    }

    const user = await this.usersService.FindById(idUser);

    const isMatchPass = await bcrypt.compare(
      updatePasswordDto.old_password,
      user.password,
    );

    if (!isMatchPass) {
      throw new HttpException(
        'old password is incorrect',
        HttpStatus.BAD_REQUEST,
      );
    }

    user.password = updatePasswordDto.new_password;
    return this.usersRepository.save(user);
  }
}
