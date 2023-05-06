import { AuthService } from './auth.service';
import { UsersService } from './../users/users.service';
import { CreateUserDto } from './../users/dto/create-user.dto';
import { User } from '../users/entities/user.entity';
import { Repository } from 'typeorm';
import { JwtTokenManagerService } from '../security/jwt-token-manager/jwt-token-manager.service';
import { JwtService } from '@nestjs/jwt';
import { Cache } from 'cache-manager';

describe('AuthService', () => {
  let userRepository: Repository<User>;
  let authService: AuthService;
  let userService: UsersService;
  let jwtTokenManagerService: JwtTokenManagerService;
  let jwtService: JwtService;
  let cacheManager: Cache;

  const user = {
    id: '0f10cff9-4b5a-4e5a-b976-56bf56efa280',
    first_name: 'user',
    last_name: '123',
    gender: 1,
    address: 'Jakarta',
    email: 'user@example.com',
    password: 'secret',
    admin: false,
    created_at: '2023-04-30 12:58:17',
    updated_at: '2023-04-30 12:58:17',
    deleted_at: null,
  };

  const payload = new CreateUserDto();
  payload.first_name = user.first_name;
  payload.last_name = user.last_name;
  payload.gender = user.gender;
  payload.address = user.address;
  payload.email = user.email;
  payload.password = user.password;
  payload.admin = user.admin;

  beforeEach(async () => {
    jwtTokenManagerService = new JwtTokenManagerService(jwtService);
    userService = new UsersService(userRepository);
    authService = new AuthService(
      userRepository,
      userService,
      jwtTokenManagerService,
      cacheManager,
    );
  });

  describe('signUp', () => {
    it('should orchestrating the sign up action correctly', async () => {
      const result = {
        id: user.id,
        first_name: user.first_name,
        last_name: user.last_name,
        email: user.email,
      };

      jest
        .spyOn(userService, 'FindByEmail')
        .mockImplementationOnce(() => Promise.resolve(null));

      jest
        .spyOn(userService, 'add')
        .mockImplementationOnce(() => Promise.resolve(user));

      const signUp = await authService.signUp(payload);

      expect(signUp).toEqual(result);
    });

    it('should orchestrating the sign up action correctly', async () => {
      const result = {
        id: user.id,
        first_name: user.first_name,
        last_name: user.last_name,
        email: user.email,
      };

      jest
        .spyOn(userService, 'FindByEmail')
        .mockImplementationOnce(() => Promise.resolve(null));

      jest
        .spyOn(userService, 'add')
        .mockImplementationOnce(() => Promise.resolve(user));

      const signUp = await authService.signUp(payload);

      expect(signUp).toEqual(result);
    });
  });
});
