import {
  HttpException,
  HttpStatus,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { UsersService } from 'src/users/users.service';
import * as bcrypt from 'bcrypt';
import { CreateUserDto } from 'src/users/dto/create-user.dto';

@Injectable()
export class AuthService {
  constructor(
    private usersService: UsersService,
    private jwtService: JwtService,
  ) {}

  async signIn(email: string, pass: string): Promise<any> {
    const user = await this.usersService.findOne(email);
    const isMatchPass = await bcrypt.compare(pass, user.password);
    if (!isMatchPass) {
      throw new UnauthorizedException();
    }

    const payload = { email: user.email, sub: user.id };

    return {
      access_token: await this.jwtService.signAsync(payload),
    };
  }
  async signUp(createUserDto: CreateUserDto): Promise<any> {
    const verifyUser = await this.usersService.findOne(createUserDto.email);

    if (verifyUser) {
      throw new HttpException('user not available', HttpStatus.BAD_REQUEST);
    }

    const users = await this.usersService.create(createUserDto);

    const result = {
      id: users.id,
      first_name: users.first_name,
      last_name: users.last_name,
      email: users.email,
    };

    return result;
  }
}
