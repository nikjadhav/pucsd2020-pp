import { Component } from '@angular/core';
import { SearchComponent } from './search/search.component';
import { UserService} from './user.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent  {
  constructor(
    private userservice:UserService,
  )
  {}
  title = 'my-app';
  change_class(){
    
    this.userservice.get_display_search();
    this.userservice.set_display_search(true);
    this.userservice.set_display_update(true);
    console.log("change");
  }
}
