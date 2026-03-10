import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
} from '@nestjs/common';
import { PreorderService } from './preorder.service';
import { CreatePreorderDto } from './dto/create-preorder.dto';
import { UpdatePreorderDto } from './dto/update-preorder.dto';
import { ValidateMongoIdPipe } from '../common/pipes/validate-mongo-id.pipe';

@Controller('preorders')
export class PreorderController {
  constructor(private readonly preorderService: PreorderService) {}

  @Post()
  create(@Body() createPreorderDto: CreatePreorderDto) {
    return this.preorderService.create(createPreorderDto);
  }

  @Get()
  findAll() {
    return this.preorderService.findAll();
  }

  @Get(':id')
  findOne(@Param('id', ValidateMongoIdPipe) id: string) {
    return this.preorderService.findOne(id);
  }

  @Patch(':id')
  update(
    @Param('id', ValidateMongoIdPipe) id: string,
    @Body() updatePreorderDto: UpdatePreorderDto,
  ) {
    return this.preorderService.update(id, updatePreorderDto);
  }

  @Delete(':id')
  remove(@Param('id', ValidateMongoIdPipe) id: string) {
    return this.preorderService.remove(id);
  }
}
