import { Injectable } from '@angular/core'
import { BehaviorSubject } from 'rxjs'
import { HttpClient } from '@angular/common/http'
import { Thing } from './api'

@Injectable({
  providedIn: 'root'
})
export class ThingService {
  // data$ provides an observable of the available things from the server
  data$ = new BehaviorSubject<Array<Thing>>(Array<Thing>())

  constructor(private http: HttpClient) {
    http.get<Array<Thing>>('/api/things').subscribe(things => this.data$.next(things))
  }
}
