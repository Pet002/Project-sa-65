
import { RoleInterface } from "./IRole";

export interface EmployeeInterface{

    ID: number,
    Name: string;
    Surname: string;
    
    RoleID: string,
    RoleInterface:RoleInterface,

}