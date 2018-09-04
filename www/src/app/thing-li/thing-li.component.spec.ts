import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ThingLiComponent } from './thing-li.component';

describe('ThingLiComponent', () => {
  let component: ThingLiComponent;
  let fixture: ComponentFixture<ThingLiComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ThingLiComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ThingLiComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
