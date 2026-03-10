import { Module } from '@nestjs/common';
import { PreOrderService } from './pre-order.service';
import { PreOrderController } from './pre-order.controller';
import { MongooseModule } from '@nestjs/mongoose';
import { Guitar, GuitarSchema } from './schemas/pre-order.schemas';

@Module({
  controllers: [PreOrderController],
  providers: [PreOrderService],
  imports: [
    MongooseModule.forFeature([{ name: Guitar.name, schema: GuitarSchema }]),
  ],
})
export class PreOrderModule {}
