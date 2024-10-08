import type Slot from './slot.ts';

export default interface Toolbox {
	id: string,
	address: string,
	slots: Slot[],
	isHovered: boolean
}