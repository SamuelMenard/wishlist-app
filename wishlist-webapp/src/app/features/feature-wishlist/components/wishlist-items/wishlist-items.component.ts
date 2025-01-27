import { ChangeDetectionStrategy, Component, input } from '@angular/core';
import { WishlistItem } from '../../services/data/models/wishlist-item.model';

@Component({
    selector: 'app-wishlist-items',
    imports: [],
    templateUrl: './wishlist-items.component.html',
    styleUrl: './wishlist-items.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class WishlistItemsComponent {
    public readonly items = input.required<WishlistItem[]>();
    public readonly loading = input.required();
    public readonly error = input.required();
}
