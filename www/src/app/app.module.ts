import { BrowserModule } from '@angular/platform-browser'
import { NgModule } from '@angular/core'
import { HttpClientModule  } from '@angular/common/http'

import { AppComponent } from './app.component'
import { ArticleComponent } from './article/article.component'
import { FilterPipe } from './filter.pipe'


@NgModule({
  declarations: [
    AppComponent,
    ArticleComponent,
    FilterPipe 
  ],
  imports: [
    BrowserModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
