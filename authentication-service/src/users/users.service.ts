import { Injectable } from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { User } from './entities/user.entity';
import { Repository } from 'typeorm';
import * as bcrypt from 'bcrypt';

@Injectable()
export class UsersService {
  constructor(
    @InjectRepository(User) private usersRepository: Repository<User>,
  ) {}

  async create(createUserDto: CreateUserDto): Promise<any> {
    const saltOrRounds = 10;
    const hashPass = await bcrypt.hash(createUserDto.password, saltOrRounds);
    createUserDto.password = hashPass;
    const newUser = this.usersRepository.create(createUserDto);
    return this.usersRepository.save(newUser);
  }

  findAll() {
    return this.usersRepository.find();
  }

  findOne(email: string) {
    return this.usersRepository.findOneBy({ email });
  }

  async update(id: string, updateUserDto: UpdateUserDto) {
    const user = await this.findOne(id);
    return this.usersRepository.save({ ...user, ...updateUserDto });
  }

  async remove(id: string) {
    const user = await this.findOne(id);
    return this.usersRepository.remove(user);
  }
}
