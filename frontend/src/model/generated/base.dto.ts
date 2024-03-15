// Code generated by tygo. DO NOT EDIT.

//////////
// source: dto.go

export interface WithFilterRequestDTO {
  filter: string;
}
export interface PaginatedRequest extends WithFilterRequestDTO {
  page: number /* int */;
  limit: number /* int */;
}
export interface PaginatedMeta {
  total: number /* int64 */;
  lastPage: number /* int32 */;
  currentPage: number /* int32 */;
  perPage: number /* int32 */;
  prev: number /* int32 */;
  next: number /* int32 */;
}
export interface Paginated<T extends any> {
  data: T[];
  meta: PaginatedMeta;
}