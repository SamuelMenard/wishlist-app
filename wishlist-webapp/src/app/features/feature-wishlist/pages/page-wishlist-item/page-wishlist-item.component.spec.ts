import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PageWishlistItemComponent } from './page-wishlist-item.component';

describe('PageWishlistItemComponent', () => {
  let component: PageWishlistItemComponent;
  let fixture: ComponentFixture<PageWishlistItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PageWishlistItemComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PageWishlistItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
