export type LinkServiceResponse = {
  created_at: string
  id: string
  link: string
  custom_bitlinks: string[]
  long_url: string
  archived: boolean
  tags: string[]
  deeplinks: unknown[]
  references: {
    group: string
  }
}
