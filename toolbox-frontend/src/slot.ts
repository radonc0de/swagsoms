export default interface Slot {
	id: string,
	address: string,
	tags: string[],
	slots: Slot[],
	isHovered: boolean
}