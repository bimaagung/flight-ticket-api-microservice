import { Module, CacheModule } from '@nestjs/common';
import { AppController } from './app.controller';
import { redisStore } from 'cache-manager-redis-yet';
import { AppService } from './app.service';
import { TypeOrmModule } from '@nestjs/typeorm';
import { ConfigModule } from '@nestjs/config';
import { UsersModule } from './users/users.module';
import { AuthModule } from './auth/auth.module';
import { JwtTokenManagerService } from './security/jwt-token-manager/jwt-token-manager.service';
import { JwtTokenManagerModule } from './security/jwt-token-manager/jwt-token-manager.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
    }),
    CacheModule.registerAsync({
      isGlobal: true,
      useFactory: async () => ({
        store: await redisStore({
          url: 'redis://@localhost:6379',
        }),
      }),
    }),
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: process.env.DB_HOST || 'localhost',
      port: parseInt(process.env.DATABASE_PORT, 10) || 3306,
      username: process.env.DB_USER,
      password: process.env.DB_PASS,
      database: process.env.DB_NAME,
      entities: ['dist/**/*.entity.js'],
      synchronize: true,
    }),
    UsersModule,
    AuthModule,
    JwtTokenManagerModule,
  ],
  controllers: [AppController],
  providers: [AppService, JwtTokenManagerService],
})
export class AppModule {}
