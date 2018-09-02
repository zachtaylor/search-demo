import { Component } from '@angular/core'
import { ThingService } from './thing.service'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  search = ''

  constructor(public thingService : ThingService) { }

  onSearch = function(event: any) {
    this.search = event.target.value
  }
}
