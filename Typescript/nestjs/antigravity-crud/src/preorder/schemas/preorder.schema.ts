import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { HydratedDocument } from 'mongoose';

export type PreorderDocument = HydratedDocument<Preorder>;

@Schema({ timestamps: true })
export class Preorder {
  @Prop({ required: true })
  customerName: string;

  @Prop({ required: true })
  guitarName: string;

  @Prop({ required: true })
  guitarBrand: string;

  @Prop({ required: true })
  guitarPrice: number;
}

export const PreorderSchema = SchemaFactory.createForClass(Preorder);
