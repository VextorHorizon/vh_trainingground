import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { Preorder, PreorderDocument } from './schemas/preorder.schema';
import { CreatePreorderDto } from './dto/create-preorder.dto';
import { UpdatePreorderDto } from './dto/update-preorder.dto';

@Injectable()
export class PreorderService {
  constructor(
    @InjectModel(Preorder.name)
    private readonly preorderModel: Model<PreorderDocument>,
  ) {}

  async create(createPreorderDto: CreatePreorderDto): Promise<Preorder> {
    const createdPreorder = new this.preorderModel(createPreorderDto);
    return createdPreorder.save();
  }

  async findAll(): Promise<Preorder[]> {
    return this.preorderModel.find().exec();
  }

  async findOne(id: string): Promise<Preorder> {
    const preorder = await this.preorderModel.findById(id).exec();
    if (!preorder) {
      throw new NotFoundException(`Preorder with ID "${id}" not found`);
    }
    return preorder;
  }

  async update(
    id: string,
    updatePreorderDto: UpdatePreorderDto,
  ): Promise<Preorder> {
    const updatedPreorder = await this.preorderModel
      .findByIdAndUpdate(id, updatePreorderDto, { new: true })
      .exec();
    if (!updatedPreorder) {
      throw new NotFoundException(`Preorder with ID "${id}" not found`);
    }
    return updatedPreorder;
  }

  async remove(id: string): Promise<Preorder> {
    const deletedPreorder = await this.preorderModel
      .findByIdAndDelete(id)
      .exec();
    if (!deletedPreorder) {
      throw new NotFoundException(`Preorder with ID "${id}" not found`);
    }
    return deletedPreorder;
  }
}
