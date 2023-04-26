import { Module } from '@nestjs/common';
import { JwtModule } from '@nestjs/jwt';
import { JwtTokenManagerService } from './jwt-token-manager.service';
import { ConfigModule } from '@nestjs/config';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      cache: true,
    }),
    JwtModule.register({
      global: true,
      signOptions: { expiresIn: process.env.ACCESS_TOKEN_AGE },
    }),
  ],
  providers: [JwtTokenManagerService],
  exports: [JwtTokenManagerService],
})
export class JwtTokenManagerModule {}
