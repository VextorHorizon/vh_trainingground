import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type GuitarProductDocument = Guitar & Document;
@Schema() 
export class Guitar{
  @Prop({ required: true })
  name: string;

  @Prop({ required: true })
  brand: string;

  @Prop({ required: true })
  guitar: string;

  @Prop({ default: 0 })
  stock: number;
}

export const GuitarSchema = SchemaFactory.createForClass(Guitar);