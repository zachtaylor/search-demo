import { Component, OnInit, Input } from '@angular/core'
import { Thing } from '../api';

@Component({
  selector: 'article',
  templateUrl: './article.component.html',
  styleUrls: ['./article.component.css']
})
export class ArticleComponent implements OnInit {
  @Input() thing : Thing

  constructor() { }

  ngOnInit() {
  }

}
