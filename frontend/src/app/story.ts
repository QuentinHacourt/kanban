export interface Story {
  id: number;
  title: string;
  description: string;
  stat: string;
  time: number;
  developer_name: string;
}

export interface StoryInput {
  title: string;
  description: string;
  time: number;
  developer_name: string;
}
