export interface Project {
  id: number;
  title: string;
  description: string;
  team_name: string;
}

export interface ProjectInput {
  title: string;
  description: string;
  team_name: string;
}
