import { Browser } from "puppeteer";
import { ComponentExample, SnapshotBuilder } from '../render-component';
export declare function componentScreenshot(browser: Browser, snapshot: ComponentExample, variation: SnapshotBuilder | null, name: string): Promise<[string, string | undefined]>;
