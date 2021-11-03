import { Component, OnInit } from '@angular/core';
import { Story } from '../story';
import { StoryService } from '../story.service';

@Component({
  selector: 'app-kanban',
  templateUrl: './kanban.component.html',
  styleUrls: ['./kanban.component.css'],
})
export class KanbanComponent implements OnInit {
  stories: Story[] = [];
  selectedStory: Story;

  constructor(private storyService: StoryService) {
    this.selectedStory = {
      id: 0,
      title: '',
      description: '',
      stat: '',
      time: 0,
      developer_name: '',
      project_name: '',
    };
  }

  ngOnInit(): void {
    this.getStories();
  }

  getStories(): void {
    this.storyService
      .getStories()
      .subscribe((stories) => (this.stories = stories));
  }

  selectStory(story: Story): void {
    this.selectedStory = story;
  }

  delete(story: Story): void {
    this.stories = this.stories.filter((h) => h !== story);
    this.storyService.deleteStory(story.id).subscribe();
  }
}
