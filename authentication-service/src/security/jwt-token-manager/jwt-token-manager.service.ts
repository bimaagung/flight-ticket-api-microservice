import { Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';

@Injectable()
export class JwtTokenManagerService {
  constructor(private jwtService: JwtService) {}

  async createAccessToken(payload: any): Promise<any> {
    return this.jwtService.signAsync(payload, {
      secret: process.env.ACCESS_TOKEN_KEY,
    });
  }

  async refreshAccessToken(payload: any): Promise<any> {
    return this.jwtService.signAsync(payload, {
      secret: process.env.REFRESH_TOKEN_KEY,
    });
  }

  async decodePayload(token: string): Promise<any> {
    return this.jwtService.decode(token);
  }

  async verifyRefreshToken(refreshToken: string) {
    try {
      await this.jwtService.verifyAsync(refreshToken, {
        secret: process.env.REFRESH_TOKEN_KEY,
      });
    } catch (error) {
      console.log(error);
      throw new UnauthorizedException();
    }
  }
}
