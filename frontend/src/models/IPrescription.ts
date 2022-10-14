import { MedicineInterface } from './IMedicine'
import { PatientInterface } from './IPatient';
import { EmployeeInterface } from './IEmployee';

export interface PrescriptionInterface {
	Prescription_ID: number,
	PH_ID: string,
	Medicine: MedicineInterface,
	Patient:PatientInterface,
	Employee:EmployeeInterface,
	Symptom:string,
	Case_Time: Date,
}


  