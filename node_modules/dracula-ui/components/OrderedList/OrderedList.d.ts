import React, { OlHTMLAttributes } from 'react';
import { ColorMap } from '../../base/colors';
import { MarginMixin, PaddingMixin } from '../../base/spacing';
export declare const orderedListColors: Partial<ColorMap>;
/**
 * OrderedList Props
 */
export interface OrderedListProps extends OlHTMLAttributes<HTMLOListElement>, PaddingMixin, MarginMixin {
    /** The Dracula UI color for the Ordered List. */
    color?: keyof typeof orderedListColors;
    /** The numbering type of the Ordered List, same as a regular `<ol>` element. */
    type?: OlHTMLAttributes<HTMLOListElement>['type'];
}
/**
 * Ordered Lists are used to display list items in an ordered way.
 */
export declare const OrderedList: React.FC<OrderedListProps>;
