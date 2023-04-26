import {
  Body,
  Controller,
  HttpCode,
  HttpStatus,
  Post,
  Put,
} from '@nestjs/common';
import { AuthService } from './auth.service';
import { CreateUserDto } from 'src/users/dto/create-user.dto';

@Controller('auth')
export class AuthController {
  constructor(private authService: AuthService) {}

  @HttpCode(HttpStatus.OK)
  @Post()
  async postAuthentication(@Body() signInDto: Record<string, any>) {
    const result = await this.authService.loginUser(
      signInDto.email,
      signInDto.password,
    );

    return {
      status: 'ok',
      message: 'success',
      data: result.access_token,
    };
  }

  @HttpCode(HttpStatus.OK)
  @Put()
  async putAuthentication(@Body() payload: Record<string, any>) {
    const accessToken = await this.authService.refreshAuthentication(
      payload.refresh_token,
    );

    return {
      status: 'ok',
      message: 'success',
      data: {
        accessToken,
      },
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
