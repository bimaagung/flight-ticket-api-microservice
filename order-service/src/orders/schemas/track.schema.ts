import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { now, Date, HydratedDocument } from 'mongoose';

export type TrackDocument = HydratedDocument<Track>;

@Schema({ timestamps: true })
export class Track {
  @Prop({ required: true })
  id: string;

  @Prop({ required: true })
  arrival: string;

  @Prop({ required: true })
  departure: string;

  @Prop({ type: Date, required: true, default: now() })
  createdAt: Date;

  @Prop({ type: Date, required: true, default: now() })
  updatedAt: Date;

  @Prop({ type: Date, default: null })
  deletedAt: Date;
}

export const TrackSchema = SchemaFactory.createForClass(Track);
