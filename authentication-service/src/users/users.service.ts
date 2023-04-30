import { HttpException, HttpStatus, Injectable } from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { User } from './entities/user.entity';
import { Repository } from 'typeorm';
import * as bcrypt from 'bcrypt';
import { UpdateUserDto } from './dto/update-user.dto';

@Injectable()
export class UsersService {
  constructor(
    @InjectRepository(User) private usersRepository: Repository<User>,
  ) {}

  // TODO: verify email not working
  async add(createUserDto: CreateUserDto): Promise<any> {
    const saltOrRounds = 10;
    const hashPass = await bcrypt.hash(createUserDto.password, saltOrRounds);
    createUserDto.password = hashPass;
    const newUser = this.usersRepository.create(createUserDto);
    return this.usersRepository.save(newUser);
  }

  findAll() {
    return this.usersRepository.find();
  }

  FindByEmail(email: string) {
    return this.usersRepository.findOneBy({ email });
  }

  async FindById(id: string) {
    const result = await this.usersRepository.findOneBy({ id });

    if (!result) {
      throw new HttpException('user not found', HttpStatus.NOT_FOUND);
    }

    return result;
  }

  async update(id: string, updateUserDto: UpdateUserDto) {
    const user = await this.FindById(id);
    return this.usersRepository.save({ ...user, ...updateUserDto });
  }

  async remove(id: string) {
    const user = await this.FindById(id);
    return this.usersRepository.remove(user);
  }
}
