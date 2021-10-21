export interface Story {
  id: number;
  title: string;
  description: string;
  stat: string;
  time: number;
}

export interface StoryInput {
  title: string;
  description: string;
  time: number;
}
