import { IsString, IsNotEmpty, IsNumber, IsPositive, MaxLength, Max } from 'class-validator';

export class CreatePreorderDto {
  @IsString()
  @IsNotEmpty()
  @MaxLength(100)
  readonly customerName: string;

  @IsString()
  @IsNotEmpty()
  @MaxLength(100)
  readonly guitarName: string;

  @IsString()
  @IsNotEmpty()
  @MaxLength(50)
  readonly guitarBrand: string;

  @IsNumber()
  @IsPositive()
  @Max(10_000_000)
  readonly guitarPrice: number;
}
