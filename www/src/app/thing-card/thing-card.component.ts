import { Component, OnInit, Input } from '@angular/core'
import { Thing } from '../api'

@Component({
  selector: 'app-thing-card',
  templateUrl: './thing-card.component.html',
  styleUrls: ['./thing-card.component.css']
})
export class ThingCardComponent implements OnInit {
  @Input() thing : Thing

  constructor() { }

  ngOnInit() {
  }

}
