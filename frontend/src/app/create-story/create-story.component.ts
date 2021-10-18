import { Component, OnInit } from '@angular/core';
import { Story } from '../story';
import { StoryService } from '../story.service';

@Component({
  selector: 'app-create-story',
  templateUrl: './create-story.component.html',
  styleUrls: ['./create-story.component.css'],
})
export class CreateStoryComponent implements OnInit {
  constructor(private storyService: StoryService) {}

  ngOnInit(): void {}

  addStory(title: string, description: string, stat: number): void {
    if (!title || !description || !stat) {
      return;
    }
    const story: Story = {
      id: 0,
      title: title,
      description: description,
      category_id: stat,
    };
    this.storyService.addStory(story).subscribe();
  }
}
