import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { Guitar, GuitarProductDocument } from './schemas/pre-order.schemas';
import { CreatePreOrderDto } from './dto/create-pre-order.dto';
import { UpdatePreOrderDto } from './dto/update-pre-order.dto';

@Injectable()
export class PreOrderService {
  constructor(
    @InjectModel(Guitar.name)
    private guitarProductmodel: Model<GuitarProductDocument>,
  ) {}

  async create(createPreOrderDto: CreatePreOrderDto): Promise<Guitar> {
    const createdProduct = new this.guitarProductmodel(createPreOrderDto);
    return createdProduct.save();
  }

  findAll() {
    return this.guitarProductmodel.find().exec();
  }

  findOne(id: string) {
    return this.guitarProductmodel.findById(id).exec();
  }

  async update(
    id: string,
    updatePreOrderDto: UpdatePreOrderDto,
  ): Promise<Guitar> {
    const result = await this.guitarProductmodel
      .findByIdAndUpdate(id, updatePreOrderDto, { new: true })
      .exec()

    if (!result) throw new NotFoundException(`Guitar ${id} not found`);

    return result;
  }

  async remove(id: string): Promise<Guitar> {
    const result = await this.guitarProductmodel.findByIdAndDelete(id).exec();

    if (!result) throw new NotFoundException(`Guitar ${id} not found`);

    return result;
  }
}
