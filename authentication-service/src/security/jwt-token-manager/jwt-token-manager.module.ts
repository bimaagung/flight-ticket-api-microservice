import { Module } from '@nestjs/common';
import { JwtModule } from '@nestjs/jwt';
import { JwtTokenManagerService } from './jwt-token-manager.service';

@Module({
  imports: [
    JwtModule.register({
      global: true,
      signOptions: { expiresIn: '900s' },
    }),
  ],
  providers: [JwtTokenManagerService],
  exports: [JwtTokenManagerService],
})
export class JwtTokenManagerModule {}
