import { Component, OnInit } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { UserService} from '../user.service';
import { Routes, RouterModule, Router } from '@angular/router';
import { FormBuilder, FormGroup, FormControl, Validators } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})  
export class SearchComponent implements OnInit {
  data:any;
  Id: string="";
  users=[];
  fname="";
  lname="";
  email="";
  contact="";
  pass="";
  
  
  constructor(
   private formBuilder:FormBuilder,
   private userservice:UserService,
   private router:Router,
   private toastr:ToastrService,
  ) { }
  ngOnInit():void{  
  }
Reset()
{
  this.Id="";
}
Search(){

  this.userservice.get_userbyid(this.Id).subscribe
     (
       (data:any[])=>
       {
         
        console.log(data);
        var stringData = '[' + JSON.stringify(data) + ']'
        var parseData = JSON.parse(stringData)
        this.users = parseData;     
       },
       (err:any[]) =>
       {

       }
      
     )
     this.userservice.set_display_search(false);
    }
    Delete(id){
      this.toastr.success("DELETE user successfully","TITLE");
      this.userservice.del_user(id).subscribe
      (
        (response)=>
        {
          
            
            console.log(response); 
        },
        (error) =>
        {
            console.log(error);
        }
       
      )
      this.router.navigate(['/']); 
    }

     
  Update(){
    console.log("id",this.Id)
    let update_data={
      "fname":this.fname,"lname":this.lname,"email":this.email,"contact":this.contact,"pass":this.pass
    }
    this.userservice.update_user(this.Id,update_data).subscribe
    (
      (data:any[])=>
      {

      },
      (err:any[]) =>
      {
        
      }
    )
    this.router.navigate(['/']); //navigate to ho
  }



    is_display_search(){
      return this.userservice.get_display_search();
   }
   /*change display settings */
   update_display_setting(){  
     this.userservice.set_display_search(true);
     this.userservice.set_display_update(false);
     /*setting up already filled values*/
     this.fname=this.users[0].fname;
     this.lname=this.users[0].lname;
     this.email=this.users[0].email;
     this.pass=this.users[0].pass;
     this.contact=this.users[0].contact;
  
   }
   is_display_update(){
     return this.userservice.get_display_update();
   }
  


}