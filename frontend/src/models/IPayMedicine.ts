import { EmployeeInterface } from "./IEmployee";
import { MedicineLabelInterface } from "./IMedicineLabel";
import { PerscriptionInterface } from "./IPrescription";

export interface PayMedicineInterface {

    ID : number,
    Amount : number,
    Price : number,
    PayDate : Date, 
    EmployeeID : number,
    Employee: EmployeeInterface,
    PerscriptionID : number,
    Perscription : PerscriptionInterface,
    MedicineLabelID : number,
    MedicineLabel : MedicineLabelInterface
    

}