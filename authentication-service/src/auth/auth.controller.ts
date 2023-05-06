import {
  Body,
  Controller,
  Delete,
  HttpCode,
  HttpStatus,
  Param,
  Post,
  Put,
} from '@nestjs/common';
import { AuthService } from './auth.service';
import { CreateUserDto } from 'src/users/dto/create-user.dto';
import { UpdatePasswordDto } from 'src/auth/dto/update-password.dto';
import { successResponse } from '../utils/response';

@Controller('auth')
export class AuthController {
  constructor(private authService: AuthService) {}

  @HttpCode(HttpStatus.OK)
  @Post()
  async postAuthentication(@Body() signInDto: Record<string, any>) {
    const { accessToken, refreshToken } = await this.authService.loginUser(
      signInDto.email,
      signInDto.password,
    );

    return successResponse({
      accessToken,
      refreshToken,
    });
  }

  @HttpCode(HttpStatus.OK)
  @Put()
  async putAuthentication(@Body() payload: Record<string, any>) {
    const accessToken = await this.authService.refreshAuthentication(
      payload.refresh_token,
    );

    return successResponse({
      accessToken,
    });
  }

  @HttpCode(HttpStatus.OK)
  @Delete()
  async deleteAuthentication(@Body() payload: Record<string, any>) {
    await this.authService.logoutUser(payload.refresh_token);

    return successResponse();
  }

  @HttpCode(HttpStatus.OK)
  @Post('register')
  async signUp(@Body() createUserDto: CreateUserDto) {
    const result = await this.authService.signUp(createUserDto);
    return successResponse(result);
  }

  @HttpCode(HttpStatus.OK)
  @Post('update/password/:id')
  async updatePassword(
    @Param('id') idUser: string,
    @Body() updatePasswordDto: UpdatePasswordDto,
  ) {
    await this.authService.updatePassword(idUser, updatePasswordDto);
    return successResponse();
  }
}
