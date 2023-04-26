import { Module } from '@nestjs/common';
import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';
import { UsersModule } from 'src/users/users.module';
import { JwtTokenManagerModule } from 'src/security/jwt-token-manager/jwt-token-manager.module';

@Module({
  imports: [UsersModule, JwtTokenManagerModule],
  controllers: [AuthController],
  providers: [AuthService],
  exports: [AuthService],
})
export class AuthModule {}
