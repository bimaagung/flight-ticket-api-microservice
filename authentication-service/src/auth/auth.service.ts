import { Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { UsersService } from 'src/users/users.service';

@Injectable()
export class AuthService {
  constructor(
    private usersService: UsersService,
    private jwtService: JwtService,
  ) {}

  async signIn(email: string, pass: string): Promise<any> {
    const user = await this.usersService.findOne(email);
    console.log(user);
    if (user?.password !== pass) {
      throw new UnauthorizedException();
    }

    const payload = { email: user.email, sub: user.id };

    return {
      access_token: await this.jwtService.signAsync(payload),
    };
  }
}
