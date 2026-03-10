import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { PreOrderModule } from './pre-order/pre-order.module';
import { ConfigModule } from '@nestjs/config';

@Module({
  imports: [
    ConfigModule.forRoot({ isGlobal: true }),
    MongooseModule.forRoot(process.env.MONGODB_URI!),
    PreOrderModule,
  ],
})
export class AppModule {}
