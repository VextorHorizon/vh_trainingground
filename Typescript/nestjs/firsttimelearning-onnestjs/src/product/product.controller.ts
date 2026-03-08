import { Controller, Get , Post, Body} from '@nestjs/common';
import { ProductService } from './product.service';

@Controller('product')
export class ProductController {
 constructor(private readonly productService: ProductService){}
    @Get()
    returnAmfine(){
        return this.productService.returnAmfine();
    }

    @Post()
    addUser(@Body() user){
        return this.productService.addUser(user);
    }

}


