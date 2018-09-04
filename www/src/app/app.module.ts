import { BrowserModule } from '@angular/platform-browser'
import { NgModule } from '@angular/core'
import { MatCardModule, } from '@angular/material/card'
import { MatInputModule, } from '@angular/material/input'
import { MatListModule, } from '@angular/material/list'
import { FlexLayoutModule } from '@angular/flex-layout'
import { HttpClientModule  } from '@angular/common/http'

import { AppComponent } from './app.component'
import { FilterPipe } from './filter.pipe'
import { ThingCardComponent } from './thing-card/thing-card.component'
import { ThingLiComponent } from './thing-li/thing-li.component'


@NgModule({
  declarations: [
    AppComponent,
    FilterPipe,
    ThingCardComponent,
    ThingLiComponent
  ],
  imports: [
    BrowserModule,
    MatCardModule,
    MatInputModule,
    MatListModule,
    FlexLayoutModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
