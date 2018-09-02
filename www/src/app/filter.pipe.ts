import { Pipe, PipeTransform } from '@angular/core'
import { Thing } from './api'

@Pipe({
  name: 'filter'
})
export class FilterPipe implements PipeTransform {
  transform(things: Array<Thing>, search: string): any[] {
    if (!things) return []
    if (!search) return things
    search = search.toLowerCase();
    return things.filter(thing => {
    return thing.Name.toLowerCase().includes(search) || thing.Data.toLowerCase().includes(search)
    })
  }
}
