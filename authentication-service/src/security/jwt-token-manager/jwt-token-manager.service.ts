import { Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { jwtConstants } from 'src/security/jwt-token-manager/constants';

@Injectable()
export class JwtTokenManagerService {
  constructor(private jwtService: JwtService) {}

  async createAccessToken(payload: any): Promise<any> {
    return this.jwtService.signAsync(payload, {
      secret: jwtConstants.accessTokenKey,
    });
  }

  async refreshAccessToken(payload: any): Promise<any> {
    return this.jwtService.signAsync(payload, {
      secret: jwtConstants.refreshTokenKey,
    });
  }

  async decodePayload(token: string): Promise<any> {
    return this.jwtService.decode(token);
  }

  async verifyRefreshToken(refreshToken: string) {
    try {
      await this.jwtService.verifyAsync(refreshToken, {
        secret: jwtConstants.refreshTokenKey,
      });
    } catch (error) {
      console.log(error);
      throw new UnauthorizedException();
    }
  }
}
