import { Component, OnInit, Input } from '@angular/core'
import { Thing } from '../api'

@Component({
  selector: 'app-thing-li',
  templateUrl: './thing-li.component.html',
  styleUrls: ['./thing-li.component.css']
})
export class ThingLiComponent implements OnInit {
  @Input() thing : Thing

  constructor() { }

  ngOnInit() {
  }

}
