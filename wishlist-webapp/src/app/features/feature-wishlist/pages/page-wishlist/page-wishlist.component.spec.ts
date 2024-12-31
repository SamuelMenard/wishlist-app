import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PageWishlistComponent } from './page-wishlist.component';

describe('PageWishlistComponent', () => {
  let component: PageWishlistComponent;
  let fixture: ComponentFixture<PageWishlistComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PageWishlistComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PageWishlistComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
