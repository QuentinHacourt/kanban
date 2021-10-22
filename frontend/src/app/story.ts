export interface Story {
  id: number;
  title: string;
  description: string;
  stat: string;
  time: number;
  userName: string;
}

export interface StoryInput {
  title: string;
  description: string;
  time: number;
  userName: string;
}
