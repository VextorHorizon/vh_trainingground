import { PartialType } from '@nestjs/mapped-types';
import { CreatePreOrderDto } from './create-pre-order.dto';

export class UpdatePreOrderDto extends PartialType(CreatePreOrderDto) {
    readonly name?: string;
    readonly brand?: string;
    readonly guitar?: string;
    readonly stock?: number;
}
