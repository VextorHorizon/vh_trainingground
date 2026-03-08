import { Injectable } from '@nestjs/common';

@Injectable()
export class ProductService {
    private readonly Users = [
        {id: 1, name: 'Vextor', At:'admin'},
        {id: 2, name: 'Copter', At:'user'}
    ];

    addUser(user){
        this.Users.push(user)
        return this.Users
    }

    returnAmfine(){
        return this.Users
    }

}
