import {
  Body,
  CACHE_MANAGER,
  Controller,
  HttpCode,
  HttpStatus,
  Inject,
  Post,
} from '@nestjs/common';
import { AuthService } from './auth.service';
import { CreateUserDto } from 'src/users/dto/create-user.dto';
import { Cache } from 'cache-manager';

@Controller('auth')
export class AuthController {
  constructor(
    private authService: AuthService,
    @Inject(CACHE_MANAGER) private cacheManager: Cache,
  ) {}

  @HttpCode(HttpStatus.OK)
  @Post('login')
  async signIn(@Body() signInDto: Record<string, any>) {
    const result = await this.authService.signIn(
      signInDto.email,
      signInDto.password,
    );
    await this.cacheManager.set(result.id, result.access_token);
    return {
      status: 'ok',
      message: 'success',
      data: result.access_token,
    };
  }

  @HttpCode(HttpStatus.OK)
  @Post('register')
  async signUp(@Body() createUserDto: CreateUserDto) {
    const result = await this.authService.signUp(createUserDto);
    return {
      status: 'ok',
      message: 'success',
      data: result,
    };
  }
}
