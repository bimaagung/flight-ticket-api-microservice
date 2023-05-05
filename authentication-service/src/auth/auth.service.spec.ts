import { Test, TestingModule } from '@nestjs/testing';
import { AuthService } from './auth.service';
import { describe } from 'node:test';

describe('AuthService', () => {
  let service: AuthService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [AuthService],
    }).compile();

    service = module.get<AuthService>(AuthService);
  });

  describe('signUp', () => {
    it('should orchestrating the sign up action correctly', () => {
      const result = {
        id: '0f10cff9-4b5a-4e5a-b976-56bf56efa280',
        first_name: 'user',
        last_name: 'active',
        email: 'user@example.com',
      };

      jest.spyOn(service, '')
    });
  });
});
