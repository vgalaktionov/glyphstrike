import { Display } from 'rot-js';

export class Engine {
    private display: Display;

    constructor() {
        this.display = new Display({
            width: 160,
            height: 100,
            layout: 'rect',
        });

        const element = this.display.getContainer()!;
        if (!(element instanceof HTMLCanvasElement)) throw new Error('invalid renderer');

        element.style.width = window.innerWidth + 'px';
        element.style.height = window.innerHeight + 'px';

        const context = element.getContext('2d')!;
        context.imageSmoothingEnabled = false;
        context.font = `16px ui-monospace,Menlo,Monaco,"Cascadia Mono","Segoe UI Mono","Roboto Mono","Oxygen Mono","Ubuntu Monospace","Source Code Pro","Fira Mono","Droid Sans Mono","Courier New", monospace`;

        document.body.appendChild(element);
    }

    public run() {
        this.display.draw(5, 4, '@', null, null);
        this.display.draw(6, 4, '@', null, null);
        this.display.draw(15, 4, '%', '#0f0', null); /* foreground color */
        this.display.draw(25, 4, '#', '#f00', '#009'); /* and background color */
    }
}
