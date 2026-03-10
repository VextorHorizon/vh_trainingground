import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { PreorderController } from './preorder.controller';
import { PreorderService } from './preorder.service';
import { Preorder, PreorderSchema } from './schemas/preorder.schema';

@Module({
  imports: [
    MongooseModule.forFeature([
      { name: Preorder.name, schema: PreorderSchema },
    ]),
  ],
  controllers: [PreorderController],
  providers: [PreorderService],
})
export class PreorderModule {}
